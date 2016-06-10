import { Injectable } from '@angular/core';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { ViewData } from './view-data.model';

@Injectable()
export class SignUpService {

  private resolveUrl = '/api/sign-up/resolve';

  resolveView(): Promise<ViewData> {
    return fetch(this.resolveUrl, {credentials: 'include'})
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

  submitEmail() {
  }
}

