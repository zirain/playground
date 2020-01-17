import { AngularTourOfHeroesPage } from './app.po';

describe('angular-tour-of-heroes App', () => {
  let page: AngularTourOfHeroesPage;

  beforeEach(() => {
    page = new AngularTourOfHeroesPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!!');
  });
});
