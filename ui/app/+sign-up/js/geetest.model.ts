export interface CaptchaResponse {
  geetestId: string;
  captchaId: string;
  mode: number;
}

export interface CaptchaRequest {
  captchaId: string;
  mode: number;
  key: string;
  hash: string;
}
