import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
import { Subscriber } from 'rxjs/Subscriber';

import { Errors } from './errors.model';
import { EmailSubmitRequest, EmailViewResponse } from './email.model';
import { ViewResponse } from './sign-up.model';

@Component({
  moduleId: module.id,
  selector: 'sign-up-email',
  templateUrl: '../tmpl/email.component.tmpl'
})
export class EmailComponent {

  Errors = Errors; // Import Errors enum

  email: string;

  emailError: Observable<number>;

  @Output() updateView = new EventEmitter();

  @Input() viewResponse: EmailViewResponse;

  private captcha: GeetestCaptcha;

  private emailRegExp = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

  private emailStream = new Subject<string>();

  private submitUrl = '/api/sign-up/submit-email';

  ngOnInit() {
    this.initSessionCookie();
    this.initGeetestCaptcha();
    this.initEmailValidation();
  }

  initSessionCookie() {
    if (this.viewResponse.cookie) {
      let sid = this.viewResponse.cookie.value;
      let date = new Date();
      date.setTime(date.getTime() + this.viewResponse.cookie.maxAge * 1000);
      let expires = date.toUTCString();
      document.cookie = `signup_sid=${sid}; Expires=${expires}`;
    }
  }

  initGeetestCaptcha() {
    System
      .import('http://static.geetest.com/static/tools/gt.js')
      .then(() => {
        let options: GeetestOptions = {
          gt: this.viewResponse.captcha.geetestId,
          challenge: this.viewResponse.captcha.captchaId,
          offline: !this.viewResponse.captcha.mode
        };
        initGeetest(options, (captcha: GeetestCaptcha) => {
          captcha.appendTo("#geetest-captcha");
          this.captcha = captcha;
        });
      });
  }

  initEmailValidation() {
    this.emailError = this.emailStream
      .distinctUntilChanged()
      .switchMap((email: string) => {
        if (!this.emailRegExp.test(email)) {
          return Observable.of(Errors.MalformedEmailAddress);
        }
        if (email === 'coolwust@gmail.com') {
          return Observable.create((subscriber: Subscriber<number>) => {
            window.setTimeout(() => {
              subscriber.next(Errors.EmailAddressExists);
              subscriber.complete();
            }, 1000);
          });
        }
        return Observable.of(null);
      });
  }

  onSubmit() {
    let validate = this.captcha.getValidate();
    if (!validate) {
      return;
    }
    let request: EmailSubmitRequest = {
      email: this.email,
      captcha: {
        captchaId: validate.geetest_challenge,
        key: validate.geetest_seccode,
        hash: validate.geetest_validate,
        mode: this.viewResponse.captcha.mode
      }
    };
    // todo mime
    let config: RequestInit = {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify(request)
    }
    fetch(this.submitUrl, config)
      .then((resp: Response) => resp.json())
      .then((resp: ViewResponse) => this.updateView.emit(resp));
  }

  onValidateEmail(email: string) {
    this.emailStream.next(email);
  }
}
