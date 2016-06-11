import { Captcha } from './geetest.model';

export interface SubmitEmailRequest {
  email: string;
  captcha: Captcha;
}
