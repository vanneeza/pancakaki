// Set your publishable key: remember to change this to your live publishable key in production
// See your keys here: https://dashboard.stripe.com/apikeys
const stripe = Stripe('pk_test_51NBe0ZIEAauHIyeskwkwtwlsGg58WpgeVmiWKuJPr1uv2LGQyBluJuad7w59iOiIj0mSCxxfS0ZiWUAdvN4GPxfe0023wgk8cr');
const options = {
    clientSecret: '{{CLIENT_SECRET}}',
    // Fully customizable with appearance API.
    appearance: {/*...*/},
  };


  (async () => {
    const response = await fetch('/secret');
    const {client_secret: clientSecret} = await response.json();
    // Call stripe.confirmCardPayment() with the client secret.
  })();
  // Set up Stripe.js and Elements to use in checkout form, passing the client secret obtained in step 3
  const elements = stripe.elements(options);
  
  // Create and mount the Payment Element
  const paymentElement = elements.create('payment');
  paymentElement.mount('#payment-element');