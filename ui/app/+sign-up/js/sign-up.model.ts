export enum Step {
  Entry,
  Confirmation,
  Detail
}

export interface View {
  step: Step;
}

export enum Alert {
  NetworkError = -100,
  ServerError,
  CorruptedData,
  UnknownError,
  SessionExpired = 1
}

export enum InputStatus {
  Validating = -100,
  Validated,
  Initial = 0,
  Empty,
  Exists,
  Malformed,
  Incorrect
}

export interface Failure {
  alerts?: Alert[]
}
