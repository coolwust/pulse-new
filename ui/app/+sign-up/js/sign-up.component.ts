import { ChangeDetectorRef, ChangeDetectionStrategy, Component, Input, NgZone } from '@angular/core';
import { OnActivate, Router, RouteSegment, RouteTree } from '@angular/router';
import { Observable } from 'rxjs/Observable';

import { ConfirmationComponent } from './confirmation.component';
import { EmailComponent } from './email.component';
import { SignUpService } from './sign-up.service';
import { ViewResponse } from './response.model';

declare var initGeetest: any;

@Component({
  moduleId: module.id,
  selector: 'app-sign-up',
  templateUrl: '../tmpl/sign-up.component.tmpl',
  providers: [SignUpService],
  directives: [ConfirmationComponent, EmailComponent]
})
export class SignUpComponent implements OnActivate {

  viewResponse: ViewResponse;

  constructor(
    private router: Router,
    private signUpService: SignUpService,
    private changeDetectorRef: ChangeDetectorRef
  ) {}

  routerOnActivate(curr: RouteSegment, prev: RouteSegment) {
    
    // By reloading or typing in the addresss bar, the `prev` will be `null`,
    // in this case, it wll redirect the user to website root.
    if (document.cookie.match('login_sid') !== null) {
      this.router.navigate(['/'], prev);
      return;
    }

    this.signUpService
      .resolveView()
      .then((resp: ViewResponse) => this.viewResponse = resp )
      .then(() => this.changeDetectorRef.detectChanges());
  }

  onUpdateView(resp: ViewResponse) {
    this.viewResponse = resp;
  }
}
