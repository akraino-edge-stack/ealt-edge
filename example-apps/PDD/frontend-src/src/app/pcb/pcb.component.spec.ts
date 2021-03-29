import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PcbComponent } from './pcb.component';

describe('PcbComponent', () => {
  let component: PcbComponent;
  let fixture: ComponentFixture<PcbComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PcbComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PcbComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
