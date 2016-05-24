fetch('/register')
  .then((resp: Response) => {
    if (resp.status !== 200) {
      return Promise.reject(new Error(response.statusText));
    }
    return Promise.resolve(resp)
  })
  .then((resp: Response) => {
    return resp.json();
  })
  .then(json: any) => {
    console.log(json);
  }
  .catch((err: string) => {
    console.error(err);
  })
