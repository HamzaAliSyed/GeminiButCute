import MainChatWindow from "./components/MainChatWindow"
import SideBar from "./components/SideBar"

function App() {
  return (
    <div className="flex w-full min-h-screen">
      <SideBar />
      <MainChatWindow />
    </div>
  )
}

export default App
