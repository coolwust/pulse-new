import { Component, Input } from '@angular/core';

import { EmailViewData, ViewResponse } from './response.model';

@Component({
  moduleId: module.id,
  selector: 'sign-up-confirmation',
  templateUrl: '../tmpl/confirmation.component.tmpl'
})
export class ConfirmationComponent {

  @Input() viewResponse: ViewResponse;
}
