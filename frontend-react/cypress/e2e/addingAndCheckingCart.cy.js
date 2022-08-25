describe('empty spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')
    cy.get('.product-add-button').first().click()
    cy.url().should('include', '/')

    cy.visit('http://localhost:3000/cart')
    cy.url().should('include', '/cart')
    cy.get('.quantity').first().should(
      "have.text",
      "Quantity: 1"
    )

    cy.get('.cart-items').find('.item-in-cart').should('have.length', 1)
  })
})