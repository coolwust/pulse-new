import { Injectable } from '@angular/core';

import { ConfirmationView } from './confirmation.model';
import { EntryFailure, EntryForm } from './entry.model';
import { Alert } from './sign-up.model';

@Injectable()
export class EntryFormService {

  private endpoint = '/api/sign-up/entry-form';

  submit(form: EntryForm): Promise<ConfirmationView | EntryFailure> {
    let headers = new Headers()
    headers.append('Content-Type', 'application/json');

    let init: RequestInit = {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify(form),
      headers: headers
    }

    return fetch(this.endpoint, init)
      .then(this.extractData)
      .catch(this.handleError);
  }

  private extractData(resp: Response): ConfirmationView | EntryFailure {
    if (resp.type === 'error') {
      return <EntryFailure> {alerts: [Alert.NetworkError]};
    }
    if (resp.status !== 200) {
      return <EntryFailure> {alerts: [Alert.ServerError]};
    }
    return resp
      .json()
      .then((obj: ConfirmationView | EntryFailure) => obj)
      .catch(() => <EntryFailure> {alerts: [Alert.CorruptedData]});
  }

  private handleError(error: any): EntryFailure {
    console.error(error);
    return <EntryFailure> {alerts: [Alert.UnknownError]};
  }
}

export class EmailExistsService {

  private endpoint = '/api/sign-up/email-exists';

  check(email: string): Promise<boolean> {
    return fetch(endpoint + `?email=${email}`)
      .then(this.extractData)
      .catch(this.handleError);
  }

  private extractData(resp: Response): boolean {
    if (resp.type === 'error') {
      return 
    }
  }
}
