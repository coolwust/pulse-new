import { ChangeDetectorRef, Component, EventEmitter, Input, OnInit, Output, NgZone } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
import { Subscriber } from 'rxjs/Subscriber';

import { ConfirmationView } from './confirmation.model';
import { EntryFailure, EntryForm, EntryView } from './entry.model';
import { EntryFormService } from './entry.service';
import { InputStatus, View } from './sign-up.model';

@Component({
  moduleId: module.id,
  selector: 'sign-up-entry',
  templateUrl: '../tmpl/entry.component.tmpl',
  providers: [EntryFormService]
})
export class EntryComponent {

  InputStatus = InputStatus; // Import enum

  email: string;

  emailStatus: InputStatus;

  private emailRegExp = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

  private emailStream = new Subject<string>();

  @Output() updateView = new EventEmitter();

  @Input() view: EntryView;

  private captcha: GeetestCaptcha;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private entryFormService: EntryFormService,
    private zone: NgZone
  ) {}

  ngOnInit() {
    this.initSession();
    this.initCaptcha();
    this.initEmailStatus();
  }

  private initSession() {
    if (this.view.hasOwnProperty('session')) {
      let sid = this.view.session.id;
      let date = new Date();
      date.setTime(this.view.session.expires * 1000);
      let expires = date.toUTCString();
      document.cookie = `signup_sid=${sid}; Expires=${expires}`;
    }
  }

  private initCaptcha() {
    System
      .import('http://static.geetest.com/static/tools/gt.js')
      .then(() => {
        let options: GeetestOptions = {
          gt: this.view.captcha.geetestId,
          challenge: this.view.captcha.captchaId,
          offline: !this.view.captcha.mode
        };
        initGeetest(options, (captcha: GeetestCaptcha) => {
          captcha.appendTo("#geetest-captcha");
          this.captcha = captcha;
        });
      });
  }

  private initEmailStatus() {
    this.emailStream
      .distinctUntilChanged()
      .switchMap((email: string) => {
        if (!this.emailRegExp.test(email)) {
          return Observable.of(InputStatus.Malformed);
        }
        if (email === 'coolwust@gmail.com') {
          return Observable
            .create((subscriber: Subscriber<InputStatus>) => {
              setTimeout(() => {
                subscriber.next(InputStatus.Exists);
                this.changeDetector.detectChanges();
                subscriber.complete();
              }, 1000);
            })
            .startWith(InputStatus.Validating);
        }
        return Observable.of(InputStatus.Validated);
      })
      .startWith(InputStatus.Initial)
      .subscribe((status: InputStatus) => this.emailStatus = status);
  }

  onSubmit() {
    let validation = this.captcha.getValidate();
    if (!validation) {
      return;
    }

    let form: EntryForm = {
      email: this.email,
      captcha: {
        captchaId: validation.geetest_challenge,
        key: validation.geetest_seccode,
        hash: validation.geetest_validate,
        mode: this.view.captcha.mode
      }
    };

    this.entryFormService
      .submit(form)
      .then((obj: ConfirmationView | EntryFailure) => {
        if (obj.hasOwnProperty('step')) {
          this.updateView.emit(obj);
        } else {
          if (obj.hasOwnProperty('email')) {
            this.emailStatus = <InputStatus> obj.email;
          }
        }
      })
      .then(() => this.changeDetector.detectChanges());
  }
}
