interface GeetestOptions {
  gt: string; // Geetest ID
  challenge: string; // Captcha ID
  offline: boolean; // Captcha mode
  product?: 'float' | 'embed' | 'popup'; // Captcha appearance. Mobile can only choose 'embed'
  lang?: 'zh-cn' | 'zh-tw' | 'en' | 'ja' | 'ko';
  sandbox?: boolean; // Default to false
  width?: string; // Mobile only
}

interface GeetestCaptcha {
  appendTo(position: Element | string, callback?: () => void): void;
  bindOn(btn: Element | string, callback: () => void): void;

  disable(): void;
  enable(): void;
  refresh(): void;

  onReady(callback: () => void): void;
  onRefresh(callback: () => void): void;
  onSuccess(callback: () => void): void;
  onFail(callback: () => void): void;
  onError(callback: () => void): void;

  getValidate(): {geetest_challenge: string, geetest_validate: string, geetest_seccode: string};
}

declare function initGeetest(options: GeetestOptions, callback: (captcha: GeetestCaptcha) => void): void;
