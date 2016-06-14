import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { ViewData } from './view.model';

@Injectable()
export class ViewResolveService {

  private endpoint = '/api/sign-up/resolve-view';

  resolve(): Promise<ViewData> {
    return fetch(this.endpoint, {credentials: 'include'})
      .then(this.extractData)
      .catch(this.handleError);
  }

  extractData(resp: Response): Promise<ViewData> {
    return resp.json();
  }

  handleError(error: any) {
    console.log(error);
    throw error.message;
  }
}

