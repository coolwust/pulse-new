import { GeetestCaptcha } from './geetest-captcha.model';

export class ViewData {
  view: 'email' | 'confirmation' | 'account';
  email: string;
  sid: string;
  captcha: GeetestCaptcha;
}
