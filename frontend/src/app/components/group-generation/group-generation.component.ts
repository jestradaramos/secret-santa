import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormControlName } from '@angular/forms';

@Component({
  selector: 'app-group-generation',
  templateUrl: './group-generation.component.html',
  styleUrls: ['./group-generation.component.css']
})
export class GroupGenerationComponent implements OnInit {
  groupForm = new FormGroup({
    gname: new FormControl(''),
    money: new FormControl(''),
    deadline: new FormControl(''),
    exchange: new FormControl('')
  })

  constructor() { }

  ngOnInit() {
  }

  submit() {
    console.warn(this.groupForm.value)
  }
}
