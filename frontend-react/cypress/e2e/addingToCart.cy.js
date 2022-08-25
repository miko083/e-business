describe('empty spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')
    cy.get('.product-add-button').first().click()
    cy.url().should('include', '/')

    cy.get('.product-add-button').eq(1).click()
    cy.url().should('include', '/')
  })
})