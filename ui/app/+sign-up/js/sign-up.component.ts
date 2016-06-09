import { ChangeDetectorRef, ChangeDetectionStrategy, Component, Input, NgZone } from '@angular/core';
import { OnActivate, Router, RouteSegment, RouteTree } from '@angular/router';
import { Observable } from 'rxjs/Observable';

import { SignUpService } from './sign-up.service';
import { ViewData } from './view-data.model';

declare var initGeetest: any;

@Component({
  moduleId: module.id,
  selector: 'app-sign-up',
  templateUrl: '../tmpl/sign-up.component.tmpl',
  providers: [SignUpService]
})
export class SignUpComponent implements OnActivate {

  view: string;

  email: string;

  constructor(private router: Router, private signUpService: SignUpService, private cd: ChangeDetectorRef) {}

  onSubmit() {
    fetch('/api/sign-up/email', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-type': 'application/json'
      },
      body: '{"email": "coolwust@gmail.com"}',
      credentials: 'include'
    }).then(resp => resp.json()).then(data => console.log(data));
  }

  routerOnActivate(curr: RouteSegment, prev: RouteSegment) {
    
    if (document.cookie.match('login_sid') !== null) {
      // If the user comes to this page by reloading or typing in the addresss
      // bar, the `prev` will be `null`, in this case, it wll redirect the user
      // to website root.
      this.router.navigate(['/'], prev);
      return;
    }

    let p1 = this.signUpService
      .getViewData()
      .then((data: ViewData) => {
        this.view = data.view;
        this.cd.detectChanges();
        return data;
      });
    let p2 = System.import('http://static.geetest.com/static/tools/gt.js');
    p1.then(data => {
      document.cookie = `signup_sid=${data.sid}`;
    });
    Promise.all([p1, p2])
      .then(([data]) => {
        initGeetest({
          gt: data.captcha.geetestId,
          challenge: data.captcha.captchaId,
          offline: !data.captcha.mode
        }, (obj: any) => {
          obj.appendTo("#geetest-captcha");
          obj.onSuccess(() => {
            obj.disable();
          });
        });
      })
      .catch((err: any) => console.log(err));

  }
}
