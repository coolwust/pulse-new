import { Captcha } from './geetest.model';

export interface ViewResponse {
  view: string;
  data: any;
}

export interface EmailViewData {
  cookie?: {
    value: string;
    maxAge: number;
  };
  captcha: Captcha;
}

export interface ConfirmationViewData {
  email: string;
}

export interface AccountViewData {
  email: string;
}
