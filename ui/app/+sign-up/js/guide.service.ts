import { Injectable } from '@angular/core';

import { Alert, Failure, View } from './sign-up.model';

@Injectable()
export class GuideService {

  private endpoint = '/api/sign-up/guide';

  guide(): Promise<Failure | View> {
    return fetch(this.endpoint, {credentials: 'include'})
      .then(this.extractData)
      .catch(this.handleError);
  }

  private extractData(resp: Response): Failure | View {
    if (resp.type === 'error') {
      return <Failure> {alerts: [Alert.NetworkError]};
    }
    if (resp.status !== 200) {
      return <Failure> {alerts: [Alert.ServerError]};
    }
    return resp
      .json()
      .then((obj: Failure | View) => obj)
      .catch(() => <Failure> {alerts: [Alert.CorruptedData]});
  }

  private handleError(error: any): Failure {
    console.error(error);
    return <Failure> {alerts: [Alert.UnknownError]};
  }
}
