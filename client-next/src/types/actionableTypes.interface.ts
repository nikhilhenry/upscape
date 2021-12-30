export interface Actionable {
  _id?: string;
  id: number;
  name: string;
  inbox: boolean;
  completed: boolean;
}

export interface ActionableUpdate {
  id?: number;
  name?: string;
  inbox?: boolean;
  completed?: boolean;
}
