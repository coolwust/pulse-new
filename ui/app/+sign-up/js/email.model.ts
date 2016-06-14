import { CaptchaRequest, CaptchaResponse } from './geetest.model';
import { ViewResponse } from './sign-up.model';

export interface EmailViewResponse extends ViewResponse {
  cookie?: {
    value: string;
    maxAge: number;
  };
  captcha: CaptchaResponse;
}

export interface EmailSubmitRequest {
  email: string;
  captcha: CaptchaRequest;
}
