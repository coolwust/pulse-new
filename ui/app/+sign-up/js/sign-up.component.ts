import { ChangeDetectorRef, Component, Input } from '@angular/core';
import { OnActivate, Router, RouteSegment, RouteTree } from '@angular/router';
import { Observable } from 'rxjs/Observable';

import { ConfirmationComponent } from './confirmation.component';
import { EmailComponent } from './email.component';
import { ViewKind, View } from './view.model';
import { ViewResolveService } from './view-resolve.service';

@Component({
  moduleId: module.id,
  selector: 'app-sign-up',
  templateUrl: '../tmpl/sign-up.component.tmpl',
  providers: [SignUpService],
  directives: [ConfirmationComponent, EmailComponent]
})
export class SignUpComponent implements OnActivate {

  ViewKind = ViewKind; // Import enum

  view: View;

  constructor(
    private router: Router,
    private viewResolveService: ViewResolveService,
    private changeDetector: ChangeDetectorRef
  ) {}

  routerOnActivate(curr: RouteSegment, prev: RouteSegment) {
    
    // By reloading or typing in the addresss bar, the `prev` will be `null`,
    // in this case, it wll redirect the user to website root.
    if (document.cookie.match('login_sid') !== null) {
      this.router.navigate(['/'], prev);
      return;
    }

    this.viewResolveService
      .resolveView()
      .then((resp: ViewResponse) => this.viewResponse = resp )
      .then(() => this.changeDetector.detectChanges());
  }

  onUpdateView(resp: ViewResponse) {
    this.viewResponse = resp;
  }
}
