import MainChatWindow from "./components/MainChatWindow"
import SideBar from "./components/SideBar"

function App() {
  return (
    <div className="flex h-screen">
      <SideBar />
      <MainChatWindow />
    </div>
  )
}

export default App
