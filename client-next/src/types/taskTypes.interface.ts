export interface Task {
  _id?: string;
  name: string;
  duration: number;
  completed: boolean;
  is_tomorrow: boolean;
  created_at?: string;
  completed_at?: string;
  highlight: boolean;
  id: number;
}

export interface TaskUpdateRequest {
  name?: string;
  duration?: string;
  completed?: boolean;
  highlight?: boolean;
  id?: number;
}
