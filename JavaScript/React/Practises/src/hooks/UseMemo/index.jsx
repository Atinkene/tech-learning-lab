import { useState, useMemo } from "react"

function ReactUseMemo() {
  const [accountBalance, setAccountBalance] = useState(0);
  const [isCalculating, setIsCalculating] = useState(false);
  
  const memoizedBalance = useMemo(() => {
    console.log("🔄 Calculating new balance...")
    setIsCalculating(true);
    
    let total = accountBalance;
    for (let i = 0; i < 1000000000; i++) {
      total += i * 0.000001;
    }
    
    setTimeout(() => setIsCalculating(false), 0);
    return total;
  }, [accountBalance]);

  return (
    <div className="flex flex-col items-center justify-center w-full h-screen space-y-4">
      {isCalculating ? (
        <div className="text-xl font-semibold text-blue-500 animate-pulse">
          Calculating...
        </div>
      ) : (
        <>
          <h1 className="text-2xl font-bold">
            Account Balance: ${memoizedBalance.toFixed(2)}
          </h1>
          <p className="text-gray-600">Raw Balance: ${accountBalance}</p>
        </>
      )}
      
      <button
        className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 cursor-pointer transition disabled:opacity-50"
        onClick={() => setAccountBalance(accountBalance + 100)}
        disabled={isCalculating}
      >
        Deposit $100
      </button>
    </div>
  )
}

export default ReactUseMemo