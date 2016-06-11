//declare let initGeetest: any;

//fetch('/geetest-init')
//  .then<Error | Response>((resp: Response) => {
//    if (resp.status !== 200) {
//      return Promise.reject<Error>(new Error(resp.statusText));
//    }
//    return Promise.resolve(resp);
//  })
//  .then((resp: Response) => {
//    return resp.json<any>();
//  })
//  .then((json: any) => {
//    initGeetest({
//      gt: json.gt,
//      challenge: json.challenge,
//      product: 'float',
//      offline: !json.success
//    }, (capcha: any) => {
//      capcha.appendTo('#geetest-captcha');
//    });
//  })
//  .catch((err: Error) => {
//    console.error(err.message);
//  });

//fetch('/search?q=foo', {
//  mode: 'same-origin'
//})
//.then<Error | Response>((resp: Response) => {
//  if (resp.status !== 200) {
//    console.log(Response);
//    return Promise.reject<Error>(new Error(resp.statusText));
//  }
//  return Promise.resolve(resp);
//})
//.then((resp: Response) => {
//  return resp.json<any>();
//})
//.then((json: any) => {
//  //initGeetest({
//  //  gt: json.gt,
//  //  challenge: json.challenge,
//  //  product: 'float',
//  //  offline: !json.success
//  //}, (capcha: any) => {
//  //  capcha.appendTo('#geetest-captcha');
//  //});
//  console.log(json);
//})
//.catch((err: Error) => {
//  console.log(err.message);
//});
//
System.import('app/search');
