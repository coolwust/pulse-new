import { Component, Input } from '@angular/core';

import { ConfirmationViewResponse } from './response.model';

@Component({
  moduleId: module.id,
  selector: 'sign-up-confirmation',
  templateUrl: '../tmpl/confirmation.component.tmpl'
})
export class ConfirmationComponent {

  @Input() viewResponse: ConfirmationViewResponse;
}
