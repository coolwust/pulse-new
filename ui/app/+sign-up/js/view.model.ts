export enum ViewKind {
  Email
  Confirmation
  Detail
}

export interface View {
  view: ViewKind;
}
