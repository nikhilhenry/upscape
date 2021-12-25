export interface Task {
  _id?: string;
  name: string;
  duration: number;
  completed: boolean;
  is_tomorrow: boolean;
  created_at?: string;
  completed_at?: string;
  highlight: boolean;
}
