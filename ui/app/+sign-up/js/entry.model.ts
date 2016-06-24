import { Captcha, UsedCaptcha } from './geetest.model';
import { Failure, InputStatus, View } from './sign-up.model';

export interface EntryView extends View {
  session?: {
    id: string;
    expires: number;
  };
  captcha: Captcha;
}

export interface EntryForm {
  email: string;
  captcha: UsedCaptcha;
}

export interface EntryFailure extends Failure {
  email?: InputStatus;
  captcha?: InputStatus;
}
