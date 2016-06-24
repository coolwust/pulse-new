import { Component, Input } from '@angular/core';

import { ConfirmationView } from './confirmation.model';

@Component({
  moduleId: module.id,
  selector: 'sign-up-confirmation',
  templateUrl: '../tmpl/confirmation.component.tmpl'
})
export class ConfirmationComponent {

  @Input() view: ConfirmationView;
}
