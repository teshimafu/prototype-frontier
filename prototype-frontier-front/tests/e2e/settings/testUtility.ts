import { PortfolioResponse } from "./model";

export const stubList: PortfolioResponse[] = [
  {
    id: 1,
    title: "test1",
    uid: "testuid",
    author: "test user",
    abstruct: "abstruct",
    readme: "# readme",
    source: "sourcesourcesourcesourcesourcesource",
    link: "linklinklinklinklinklinklink",
    access_count: 10,
    created_at: "2021-01-01T12:34:56.123456Z",
    updated_at: "2021-01-01T12:34:56.123456Z"
  },
  {
    id: 2,
    title: "test2",
    uid: "testuid",
    author: "test user",
    abstruct: "abstruct",
    readme: "# readme",
    source: "sourcesourcesourcesourcesourcesource",
    link: "linklinklinklinklinklinklink",
    access_count: 20,
    created_at: "2021-01-01T12:34:56.123456Z",
    updated_at: "2021-01-01T12:34:56.123456Z"
  },
  {
    id: 3,
    title: "test3",
    uid: "testuid3",
    author: "test user3",
    abstruct: "abstruct",
    readme: "# readme",
    source: "sourcesourcesourcesourcesourcesource",
    link: "linklinklinklinklinklinklink",
    access_count: 30,
    created_at: "2021-01-01T12:34:56.123456Z",
    updated_at: "2021-01-01T12:34:56.123456Z"
  }
];

export const initStubAPI = (cy: Cypress.cy & EventEmitter) => {
  cy.intercept(
    {
      method: "GET",
      url: /\/api\/portfolios$/
    },
    {
      body: stubList
    }
  );
  cy.intercept(
    {
      method: "GET",
      url: /\/api\/portfolios\?./
    },
    req => {
      req.reply({
        body: stubList.filter(
          p =>
            p.title.includes(req.url.split("title=")[1].split("&")[0]) ||
            p.author.includes(req.url.split("author=")[1].split("&")[0]) ||
            p.readme.includes(req.url.split("readme=")[1])
        )
      });
    }
  );
  cy.intercept(
    {
      method: "GET",
      url: /\/api\/portfolios\/./
    },
    req => {
      req.reply({
        body: stubList.find(p => p.id == +req.url.split("/api/portfolios/")[1])
      });
    }
  );
};

export class TestUtility {
  public static waitTime = 100 as const;
  public static assertPortfolioDetail = (
    e: HTMLElement,
    expected: PortfolioResponse
  ) => {
    const keys = ["title", "date", "author", "link", "source"];
    expected["title"];
    keys.forEach(className => {
      switch (className) {
        case "date": {
          const date = e.getElementsByClassName(className)[0].textContent;
          assert.equal(
            date ? date.substring(0, 10) : "",
            expected.created_at.substring(0, 10)
          );
          assert.equal(
            date ? date.substring(11, 19) : "",
            expected.created_at.substring(11, 19)
          );
          break;
        }
        default: {
          assert.equal(
            e.getElementsByClassName(className)[0].textContent,
            "" + expected[className as keyof PortfolioResponse]
          );
          break;
        }
      }
    });
  };
}
