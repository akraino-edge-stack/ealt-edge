import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DataMonitorComponent } from './data-monitor.component';

describe('DataMonitorComponent', () => {
  let component: DataMonitorComponent;
  let fixture: ComponentFixture<DataMonitorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DataMonitorComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DataMonitorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
