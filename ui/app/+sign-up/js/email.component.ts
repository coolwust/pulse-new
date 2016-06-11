import { Component, Input, OnInit } from '@angular/core';

import { EmailViewData, ViewResponse } from './response.model';
import { Captcha } from './geetest.model';


declare let initGeetest: any;

@Component({
  moduleId: module.id,
  selector: 'app-sign-up-email',
  templateUrl: '../tmpl/email.component.tmpl'
})
export class EmailComponent {

  @Input() viewResponse: ViewResponse;

  email: string;

  captcha: Captcha;

  captchaObj: any;

  ngOnInit() {
    let data = <EmailViewData> this.viewResponse.data;
    if (data.cookie) {
      let date = new Date();
      date.setTime(date.getTime() + data.cookie.maxAge * 1000);
      let expires = date.toUTCString();
      document.cookie = `signup_sid=${data.cookie.value}; Path=/sign-up; Expires=${expires}`;
    }

    System
      .import('http://static.geetest.com/static/tools/gt.js')
      .then(() => {
        let config = {
          gt: data.captcha.geetestId,
          challenge: data.captcha.captchaId,
          offline: !data.captcha.mode
        };
        console.log(data.captcha);
        initGeetest(config, (obj: any) => {
          obj.appendTo("#geetest-captcha");
          this.captchaObj = obj;
        });
        this.captcha = data.captcha;
      });
  }

  onSubmit() {
    let validate = this.captchaObj.getValidate();
    if (!validate) {
      return;
    }
    let captcha = {
      mode: this.captcha.mode,
      captchaId: validate.geetest_challenge,
      key: validate.geetest_seccode,
      hash: validate.geetest_validate
    }
    console.log(captcha);
  }
}
