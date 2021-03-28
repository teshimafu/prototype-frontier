import { assert } from "chai";
import { PortfolioResponse } from "../settings/model";
import { initStubAPI, stubList, TestUtility } from "../settings/testUtility";

describe("portfolio list test", () => {
  beforeEach(() => {
    initStubAPI(cy);
  });
  it("一覧画面が正しく表示されること", () => {
    cy.visit("/").wait(TestUtility.waitTime);
    cy.get(".portfolio-link").then(element => {
      const elements = element.toArray();
      assert.equal(elements.length, stubList.length);
      elements.forEach((e, i) => {
        assertPortfolioElement(e, stubList[i]);
      });
    });
  });

  it("一覧画面で検索機能が正しく動作すること", () => {
    const searchText = "test1";
    const expectedList = stubList.filter(p => p.title === searchText);
    cy.visit("/");
    cy.get("input").type(searchText);

    //検索テスト
    cy.get(".search-unit > .btn")
      .then(element => {
        element.toArray()[0].click();
      })
      .wait(TestUtility.waitTime);
    cy.get(".portfolio-link").then(element => {
      const elements = element.toArray();
      assert.equal(elements.length, expectedList.length);
      elements.forEach((e, i) => {
        assertPortfolioElement(e, expectedList[i]);
      });
    });

    //リセットテスト
    cy.get(".search-unit > .btn")
      .then(element => {
        element.toArray()[1].click();
      })
      .wait(TestUtility.waitTime);
    cy.get(".search-unit > .form-control").should("have.value", "");

    cy.get(".portfolio-link").then(element => {
      const elements = element.toArray();
      assert.equal(elements.length, stubList.length);
      elements.forEach((e, i) => {
        assertPortfolioElement(e, stubList[i]);
      });
    });
  });

  it("一覧画面で詳細画面への遷移が正しく動作すること", () => {
    stubList.forEach((portfolio, index) => {
      cy.visit("/");
      cy.get(".portfolio-link").then(element => {
        const elements = element.toArray();
        elements[index].click();
        cy.url().should("include", `/detail/${portfolio.id}`);
        cy.get(".preview").then(element => {
          TestUtility.assertPortfolioDetail(element[0], portfolio);
        });
      });
    });
  });

  const assertPortfolioElement = (
    e: HTMLElement,
    expected: PortfolioResponse
  ) => {
    assert.equal(e.getElementsByTagName("th")[0].innerText, "" + expected.id);
    assert.equal(e.getElementsByTagName("td")[0].innerText, expected.title);
    assert.equal(e.getElementsByTagName("td")[1].innerText, expected.author);
    const dateView = e.getElementsByTagName("td")[2].innerText;
    assert.equal(
      dateView.substring(0, 10),
      expected.created_at.substring(0, 10)
    );
    assert.equal(
      dateView.substring(11, 19),
      expected.created_at.substring(11, 19)
    );
  };
});
