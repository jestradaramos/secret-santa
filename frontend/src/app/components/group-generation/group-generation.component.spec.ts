import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupGenerationComponent } from './group-generation.component';

describe('GroupGenerationComponent', () => {
  let component: GroupGenerationComponent;
  let fixture: ComponentFixture<GroupGenerationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupGenerationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupGenerationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
