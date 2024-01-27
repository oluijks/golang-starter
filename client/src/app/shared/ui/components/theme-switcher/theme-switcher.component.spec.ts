import {ComponentFixture, TestBed} from '@angular/core/testing'

import {DOCUMENT, WINDOW} from '../../../../app.tokens'
import {ThemeSwitcherComponent} from './theme-switcher.component'

describe('ThemeSwitcherComponent', () => {
  let component: ThemeSwitcherComponent
  let fixture: ComponentFixture<ThemeSwitcherComponent>

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThemeSwitcherComponent],
      providers: [
        {
          provide: WINDOW,
          useFactory: () => window,
        },
        {
          provide: DOCUMENT,
          useFactory: () => document,
        },
      ],
    }).compileComponents()

    fixture = TestBed.createComponent(ThemeSwitcherComponent)
    component = fixture.componentInstance
    fixture.detectChanges()
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })
})
