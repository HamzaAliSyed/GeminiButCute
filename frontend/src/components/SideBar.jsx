import SearchBar from "./SearchBar";
import { SiNodemon } from "react-icons/si";
import { FaChevronLeft } from "react-icons/fa";



const SideBar = () => {
    return (
            <div className="basis-[23%] flex-shrink-0 bg-slate-500 p-4">
                <div className="flex items-center w-full">
                    <SiNodemon className="text-4xl text-gray-200" />
                    <h1 className="flex-1 text-3xl font-bold text-gray-200 text-center">Chats</h1>
                    <button className="p-2 text-2xl rounded-full bg-slate-500 text-gray-200 
                           hover:bg-slate-400 transition-colors" ><FaChevronLeft /></button>
                </div>
                <SearchBar />
                <hr></hr>
            </div>
    );
};

export default SideBar;