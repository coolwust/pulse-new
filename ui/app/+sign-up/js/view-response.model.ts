import { GeetestCaptcha } from './geetest-captcha.model';

export interface ViewResponse {
  view: string;
  data: any;
}

export interface EmailViewResponse {
  view: 'email';
  data: {
    cookie?: {
      name: string;
      value: string;
      path: string;
      maxAge: number;
    };
    captcha: {
      geetestId: string;
      captchaId: string;
      mode: number;
    };
  }
}

export interface ConfirmationViewData {
  view: 'confirmation';
  email: string;
}

export interface AccountViewData {
  view: 'account';
  email: string;
}
