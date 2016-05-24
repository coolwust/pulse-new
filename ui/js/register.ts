fetch('./foo/bar').then((resp: Response) => {
  if (resp.status !== 200) {
    console.error('Something went wrong, status code:' + resp.status);
    return;
  }
  console.log(resp.headers.get('Content-Type'));
}).catch((resp: Response) => {
    console.error('Something went wrong, status code:' + resp.status);
})
