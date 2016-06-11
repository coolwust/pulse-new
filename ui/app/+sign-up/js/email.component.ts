import { Component, Input, OnInit } from '@angular/core';

import { EmailViewResponse } from './view-response.model';

declare let initGeetest: any;

@Component({
  moduleId: module.id,
  selector: 'app-sign-up-email',
  templateUrl: '../tmpl/email.component.tmpl'
})
export class EmailComponent {

  @Input() viewResponse: EmailViewResponse;

  ngOnInit() {
    System
      .import('http://static.geetest.com/static/tools/gt.js')
      .then(() => {
        let config = {
          gt: this.viewResponse.data.captcha.geetestId,
          challenge: this.viewResponse.data.captcha.captchaId,
          offline: !this.viewResponse.data.captcha.mode
        };
        initGeetest(config, (obj: any) => obj.appendTo("#geetest-captcha"));
      });
  }
}
