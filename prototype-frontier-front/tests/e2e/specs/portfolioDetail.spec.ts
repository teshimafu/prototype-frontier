import { initStubAPI, stubList, TestUtility } from "../settings/testUtility";

describe("portfolio detail test", () => {
  beforeEach(() => {
    initStubAPI(cy);
  });
  it("直接アクセスで詳細画面が表示されること", () => {
    stubList.forEach(p => {
      cy.visit("/detail/" + p.id).wait(TestUtility.waitTime);
      cy.get(".preview").then(element => {
        TestUtility.assertPortfolioDetail(element[0], p);
      });
    });
  });
});
