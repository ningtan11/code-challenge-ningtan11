import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import CurrencyInput from "react-currency-input-field";
import "./Form.css";
import exchangerate from "../../utils/utils";

function Form() {
  const [amount, setAmount] = useState(0);
  const [fromCurrency, setFromCurrency] = useState("USD");
  const [toCurrency, setToCurrency] = useState("ETH");
  const [output, setOutput] = useState(0); 

  useEffect(() => {

  }, []);

  const calculateOutput = async () => {
    const fromCurrencyRate = exchangerate.find(
      (currency) => currency.currency === fromCurrency
    ).price;
    const toCurrencyRate = exchangerate.find(
      (currency) => currency.currency === toCurrency
    ).price;
    const result = (amount / fromCurrencyRate) * toCurrencyRate;
    setOutput(result);
  };

  let navigate = useNavigate(); 
  const backToHomepage = () =>{ 
    navigate(0)
  }

  const swap = () => {
    if (document.getElementById("amount").value === "0") {
      alert("Please enter an amount to swap!");
    } else if (output === 0) {
      alert("Please calculate the amount you will get first!");
    } else {
      alert("Swap is successful!");
      // return to home page
      backToHomepage();
    }
  }

  return (
    <div>
      <h1 id="form-title">Swap Form</h1>
      <div style={{ display: "flex" }}>

        <div>
          <label>You give:</label>
          <CurrencyInput
            value={amount}
            onValueChange={(amount) => {
              setAmount(amount)}}
            allowDecimals={true}
            decimalsLimit={10000}
            allowNegativeValue={false}
            defaultValue={0}
            id="amount"
          />
        </div>

        <div className="input-from">
          <label>In:</label>
          <select
            id="from"
            value={fromCurrency}
            onChange={(e) =>
              setFromCurrency(e.target.value)
              }
          >
            {exchangerate.map((currency, index) => (
                <option key={index} value={currency.currency}>
                  {currency.currency}
                </option>
              ))
            } : (
              <option defaultValue>USD</option>
            )
          </select>
        </div>

      </div>

      <div style={{ display: "flex" }}>
        <button className="btn" onClick={() => calculateOutput()}>
          Calculate
        </button>
        <button className="btn btn-swap" onClick={() => swap()}>
          Swap
        </button>
      </div>

      <div style={{ display: "flex" }}>
        <div className="output">
          <label>You get:</label>
          <input type="text" value={output} readonly />
        </div>

        <div className="input-to">
          <label>In:</label>
          <select
            id="to"
            value={toCurrency}
            onChange={(e) => setToCurrency(e.target.value)}
          >
            {exchangerate.map((currency, index) => (
                <option key={index} value={currency.currency}>
                  {currency.currency}
                </option>
              ))
            } : (
              <option defaultValue>ETH</option>
            )
          </select>
        </div>
      </div> 

      
      
    </div>
  );
}

export default Form;
