describe("Header test", () => {
  it("headerの表示が正しいこと", () => {
    cy.visit("/");
    cy.contains("h1", "Prototype Frontier");
    cy.get("#signin").should("have.text", "サインイン");
  });
});
