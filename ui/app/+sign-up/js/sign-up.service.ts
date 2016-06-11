import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { ViewResponse } from './response.model';

@Injectable()
export class SignUpService {

  private resolveUrl = '/api/sign-up/resolve-view';

  resolveView(): Promise<ViewResponse> {
    return fetch(this.resolveUrl, {credentials: 'include'})
      .then(this.extractData)
      .catch(this.handleError);
  }

  extractData(resp: Response): Promise<ViewResponse> {
    return resp.json();
  }

  handleError(error: any) {
    console.log(error);
    throw error.message;
  }

  submitEmail() {
  }
}

