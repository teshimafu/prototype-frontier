// https://docs.cypress.io/api/introduction/api.html

describe("Header test", () => {
  it("headerの表示が正しいこと", () => {
    cy.visit("/");
    cy.get("#copylight").should("have.text", "created by teshima.fu@gmail.com");
  });
});
