import { Component, OnInit } from '@angular/core';
import * as uuid from 'uuid';

const DAYS: Day[] = [
  {
    value: 'monday',
    name: 'Monday'
  },
  {
    value: 'tuesday',
    name: 'Tuesday'
  },
  {
    value: 'wednesday',
    name: 'Wednesday'
  },
  {
    value: 'thursday',
    name: 'Thursday'
  },
  {
    value: 'friday',
    name: 'Friday'
  },
  {
    value: 'saturday',
    name: 'Saturday'
  },
  {
    value: 'sunday',
    name: 'Sunday'
  }
];

const TEMPS: number[] =
[
  60,
  61,
  62
];

@Component({
  selector: 'app-timepage',
  templateUrl: './timepage.component.html',
  styleUrls: ['./timepage.component.scss']
})
export class TimepageComponent implements OnInit {
  public days = DAYS;
  public temps = TEMPS;
  public keyframes: Keyframe[] = [];

  constructor() { }

  public onAdd(): void {
    this.keyframes.push({id: uuid.v4()} as Keyframe);
  }

  public onDelete(id: string) {
    for (let i = 0; i < this.keyframes.length; i++) {
      if (this.keyframes[i].id === id) {
        this.keyframes.splice(i, 1);
        // console.log('deleted ' + id);
        return;
      }
    }
  }

  public onSave(id: string) {
    for (const frame of this.keyframes) {
      if (frame.id === id) {
        const body = frame;

        console.log('save');
        console.log(JSON.stringify(body));
        return;
      }
    }
  }

  ngOnInit() {

  }

}
