import { Component } from '@angular/core';
import { OnActivate, Router, RouteSegment, RouteTree } from '@angular/router';

@Component({
  moduleId: module.id,
  selector: 'app-sign-up',
  templateUrl: '../tmpl/sign-up.component.tmpl'
})
export class SignUpComponent implements OnActivate {

  step: "emailAddr | checkEmail | accountInfo";

  constructor(private router: Router) {}

  routerOnActivate(curr: RouteSegment, prev: RouteSegment) {
    // TODO: Use service
    if (window.localStorage.getItem('sid_login') !== null) {
      // If the user comes to this page by reloading or typing in the addresss
      // bar, the `prev` will be `null`, in this case, it wll redirect the user
      // to website root.
      this.router.navigate(['/'], prev);
      return;
    }
    
    if (window.localStorage.getItem('sid_register') === null) {
      System.import('http://static.geetest.com/static/tools/gt.js').then(() => console.log('send requrest!')).catch(() => console.log('failed!'));
    }
  }
}
