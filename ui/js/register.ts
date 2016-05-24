fetch('/geetest-init')
  .then<Error | Response>((resp: Response) => {
    if (resp.status !== 200) {
      return Promise.reject<Error>(new Error(resp.statusText));
    }
    return Promise.resolve(resp);
  })
  .then((resp: Response) => {
    return resp.json<any>();
  })
  .then((json: any) => {
    console.log(json);
  })
  .catch((err: Error) => {
    console.error(err.message);
  })
