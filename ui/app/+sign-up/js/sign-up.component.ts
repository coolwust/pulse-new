import { ChangeDetectorRef, Component } from '@angular/core';
import { OnActivate, Router, RouteSegment } from '@angular/router';

import { ConfirmationComponent } from './confirmation.component';
import { EntryComponent } from './entry.component';
import { GuideService } from './guide.service';
import { Failure, Step, View } from './sign-up.model';

@Component({
  moduleId: module.id,
  selector: 'app-sign-up',
  templateUrl: '../tmpl/sign-up.component.tmpl',
  providers: [GuideService],
  directives: [ConfirmationComponent, EntryComponent]
})
export class SignUpComponent implements OnActivate {

  Step = Step; // Import enum

  view: View;

  constructor(
    private router: Router,
    private guideService: GuideService,
    private changeDetector: ChangeDetectorRef
  ) {}

  routerOnActivate(curr: RouteSegment, prev: RouteSegment) {
    if (this.redirectLogin(prev)) {
      return;
    }
    this.initView();
  }

  // By reloading or typing in the addresss bar, the `prev` will be
  // `null`, in this case, it wll redirect the user to website root.
  private redirectLogin(prev: RouteSegment): boolean {
    if (document.cookie.match('login_sid') !== null) {
      this.router.navigate(['/'], prev);
      return true;
    }
    return false;
  }

  private initView() {
    this.guideService
      .guide()
      .then((obj: Failure | View) => {
        if (obj.hasOwnProperty('step')) {
          this.view = <View> obj;
        } else {
          // TODO: Handle Alerts
          console.log(obj);
        }
      })
      .then(() => this.changeDetector.detectChanges());
  }
}
