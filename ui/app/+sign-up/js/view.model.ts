import { GeetestCaptcha } from './geetest-captcha.model';

export interface ViewResponse {
  view: 'email' | 'confirmation' | 'account';
  data: EmailViewData | ConfirmationData | 
}

export interface EmailViewData {
}
