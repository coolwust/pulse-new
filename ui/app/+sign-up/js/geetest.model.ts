export interface Captcha {
  geetestId: string;
  captchaId: string;
  mode: number;
}

export interface UsedCaptcha {
  captchaId: string;
  mode: number;
  key: string;
  hash: string;
}
