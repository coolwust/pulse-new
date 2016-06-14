import { ViewResponse } from './sign-up.model';

export interface ConfirmationViewResponse extends ViewResponse {
  email: string;
}
