import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

export class SignUpService {

  private sessionUrl = '/sign-up/session';

  constructor(private http: Http) {}

  //getSessionId(): Observable<string> {
  //  return this.http
  //    .get(sessionUrl)
  //    .map(extractSessionId)
  //    .catch(handleError);
  //}

  //private extractSessionId(res: Response) {
  //}
}

