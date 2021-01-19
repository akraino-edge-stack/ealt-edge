import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DataFetchComponent } from './data-fetch.component';

describe('DataFetchComponent', () => {
  let component: DataFetchComponent;
  let fixture: ComponentFixture<DataFetchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DataFetchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DataFetchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
