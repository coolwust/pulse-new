import { Component, OnInit } from '@angular/core';
import { HTTP_PROVIDERS } from '@angular/http';
import { ROUTER_DIRECTIVES, ROUTER_PROVIDERS, Routes, Route } from '@angular/router';

import { IndexComponent } from './index.component';
import { NotFoundComponent } from './not-found.component';
import { SignUpComponent } from '../+sign-up/js/sign-up.component';

@Component({
  moduleId: module.id,
  selector: 'app',
  templateUrl: '../tmpl/app.component.tmpl',
  directives: [ROUTER_DIRECTIVES],
  providers: [HTTP_PROVIDERS, ROUTER_PROVIDERS]
})
@Routes([
  new Route({path: '/', component: IndexComponent}),
  new Route({path: '/register', component: SignUpComponent}),
  new Route({path: '*', component: NotFoundComponent})
])
export class AppComponent {}
