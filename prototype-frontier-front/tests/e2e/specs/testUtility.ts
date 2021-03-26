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
    createdAt: "2021-01-01T12:34:56.123456Z",
    updatedAt: "2021-01-01T12:34:56.123456Z"
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
    createdAt: "2021-01-01T12:34:56.123456Z",
    updatedAt: "2021-01-01T12:34:56.123456Z"
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
    createdAt: "2021-01-01T12:34:56.123456Z",
    updatedAt: "2021-01-01T12:34:56.123456Z"
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
}
